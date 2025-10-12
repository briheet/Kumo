{
  description = "Kumo flake for NATS cluster development shell and services";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    systems.url = "github:nix-systems/default";
    process-compose-flake.url = "github:Platonic-Systems/process-compose-flake";
    services-flake.url = "github:juspay/services-flake";
  };

  outputs = inputs:
    inputs.flake-parts.lib.mkFlake { inherit inputs; } {
      systems = import inputs.systems;
      imports = [ inputs.process-compose-flake.flakeModule ];

      perSystem = { pkgs, config, lib, ... }: {
        process-compose."default" = { config, ... }:
          let
            commonClusterSettings = {
              accounts."$SYS".users = [
                { user = "admin"; pass = "admin"; }
              ];
              accounts.js = {
                jetstream = "enabled";
                users = [
                  { user = "js"; pass = "js"; }
                ];
              };
              jetstream.max_file = "128M";
              cluster.name = "default-cluster";
              cluster.routes = [
                "nats://localhost:14248"
                "nats://localhost:24248"
                "nats://localhost:34248"
              ];
            };
          in {
            imports = [ inputs.services-flake.processComposeModules.default ];

            services.nats-server."nats".enable = true;

            services.nats-server."nats-1" = {
              enable = true;
              settings = lib.recursiveUpdate commonClusterSettings {
                port = 14222;
                monitor_port = 18222;
                cluster.port = 14248;
              };
            };

            services.nats-server."nats-2" = {
              enable = true;
              settings = lib.recursiveUpdate commonClusterSettings {
                port = 24222;
                monitor_port = 28222;
                cluster.port = 24248;
              };
            };

            services.nats-server."nats-3" = {
              enable = true;
              settings = lib.recursiveUpdate commonClusterSettings {
                port = 34222;
                monitor_port = 38222;
                cluster.port = 34248;
              };
            };

            settings.processes.test = {
              command = pkgs.writeShellApplication {
                name = "nats-test";
                runtimeInputs = [ pkgs.natscli ];
                text = ''
                  # standalone
                  nats account info -s nats://localhost:4222

                  # nats cluster with jetstream enabled
                  nats server info -s nats://admin:admin@localhost:14222

                  export NATS_URL=nats://js:js@localhost:14222
                  nats account info

                  # https://docs.nats.io/nats-concepts/jetstream/js_walkthrough
                  nats stream add my_stream --subjects=foo --storage=memory --replicas=3 --defaults
                  nats stream info my_stream
                  nats pub foo --count=5 "publication #{{Count}} @ {{TimeStamp}}"
                  nats consumer add my_stream pull_consumer --pull --replicas=3 --defaults
                  nats consumer next my_stream pull_consumer --count 5
                  nats stream rm -f my_stream
                '';
              };
              depends_on = {
                "nats".condition = "process_healthy";
                "nats-1".condition = "process_healthy";
                "nats-2".condition = "process_healthy";
                "nats-3".condition = "process_healthy";
              };
            };
          };

        devShells.default = pkgs.mkShell {
          inputsFrom = [
            config.process-compose."default".services.outputs.devShell
          ];
          nativeBuildInputs = [
            pkgs.go_1_25
            pkgs.gopls
            pkgs.golangci-lint-langserver
            pkgs.delve
          ];
        };
      };
    };
}
