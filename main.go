package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/common/workspace"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	fmt.Println("starting run...")

	ctx := context.Background()

	stackName := "some-stack"
	config := auto.ConfigMap{
		"foo": auto.ConfigValue{
			Value:  "foo",
			Secret: false,
		},
		"bar": auto.ConfigValue{
			Value:  "bar",
			Secret: true, // Secret values cannot be decrypted, bc there's no encryption salt.
		},
	}

	// Remove the config file & pulumi state, to get us to a zero state.
	_ = os.Remove("./program/Pulumi.some-stack.yaml")
	_ = os.Remove("./program/Pulumi.yaml")
	_ = os.RemoveAll(".pulumi")

	fmt.Println("pulumi files removed...")

	// Get a workspace and run an initial update, so that some state exists
	w := newWorkspace(ctx)

	s, err := auto.UpsertStack(ctx, stackName, w)
	if err != nil {
		return err
	}

	if err = s.SetAllConfig(ctx, config); err != nil {
		return err
	}

	_, err = s.Up(ctx)
	if err != nil {
		return err
	}

	// Remove the stack config and reinstantiate the stack and the workspace
	_ = os.Remove("./program/Pulumi.some-stack.yaml")
	w = newWorkspace(ctx)
	log.Printf("Pulumi version: %s", w.PulumiVersion())

	s, err = auto.UpsertStack(ctx, stackName, w)
	if err != nil {
		return err
	}

	// Now set all config and try to read it back. This should fail.
	if err = s.SetAllConfig(ctx, config); err != nil {
		return err
	}

	// Error here
	_, err = s.GetAllConfig(ctx)
	if err != nil {
		return fmt.Errorf("GetAllConfig: %w", err)
	}

	panic("did not reproduce")
}

func newWorkspace(ctx context.Context) auto.Workspace {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	w, err := auto.NewLocalWorkspace(ctx,
		auto.EnvVars(map[string]string{
			"PULUMI_CONFIG_PASSPHRASE": "asdf",
		}),
		auto.Project(workspace.Project{
			Name:    "auto-decrypt",
			Runtime: workspace.NewProjectRuntimeInfo("go", nil),
			Backend: &workspace.ProjectBackend{
				URL: fmt.Sprintf("file://%s", cwd),
			},
		}),
		auto.WorkDir("program"),
		auto.SecretsProvider("passphrase"),
	)

	if err != nil {
		panic(err)
	}
	return w
}
