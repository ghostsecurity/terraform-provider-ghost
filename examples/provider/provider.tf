terraform {
  required_providers {
    ghost = {
      source  = "ghostsecurity/ghost"
      version = "~> 0.1.0"
    }
  }
}

# Configure the provider using the GHOST_API_KEY from the environment
provider "ghost" {
}
