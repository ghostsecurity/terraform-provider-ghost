# terraform-provider-ghost
The [Ghost Security](https://ghostsecurity.com/) terraform provider.

## Development
Refer to the [Hashicorp Documentation](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider) for more details on developing custom providers.

### Pre-requisites
- [Terraform](https://developer.hashicorp.com/terraform/install) >= 1.0
- [Go](https://go.dev/learn/) >= 1.23
- [golangci-lint](https://github.com/golangci/golangci-lint) >= 1.61

### Install Provider
To build and install the provider locally run `make install`.

To configure terraform to use the locally installed provider instead of the terraform registry, edit the `~/.terraformrc` file as follows:

```hcl
provider_installation {
  dev_overrides {
      "ghostsecurity/ghost" = "<PATH>"
  }

  direct {}
}
```

The value for `<PATH>` should be `GOBIN` path where Go installs binaries. This will either be the value returned from `go env GOBIN` or the default path of `$HOME/go/bin` if empty.

### Use Local Provider
With the terraform provider installed locally you should be abled to use it by specifying the `source` in the `required_providers` block to match what was specified in the `dev_overrides` block of your `~/.terraformrc`.

Example `main.tf`:
```hcl
terraform {
  required_providers {
    ghost = {
      source = "ghostsecurity/ghost",
    }
  }
}

provider "ghost" {
}

resource "ghost_log_forwarder" "example" {
    name = "example"
}

output "subject_id" {
    value = ghost_log_forwarder.example.subject_id
}
```

### Acceptance Tests
To run the acceptance tests you must first set a valid `GHOST_API_KEY` in the environment.
Running `make test-acc` will run the acceptance tests to exercise the various terraform
resources using the actual API which will create/destroy _real resources_ in the platform.
