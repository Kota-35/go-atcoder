{
  description = "go-atcoder";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };

  outputs =
    { nixpkgs, ... }:
    let
      supportedSystems = [
        "aarch64-darwin"
        "x86_64-darwin"
        "x86_64-linux"
        "aarch64-linux"
      ];
      forEachSystem =
        f: nixpkgs.lib.genAttrs supportedSystems (system: f nixpkgs.legacyPackages.${system});
    in
    {
      formatter = forEachSystem (pkgs: pkgs.nixfmt-tree);

      devShells = forEachSystem (pkgs: {
        default =
          let
            # online-judge-api-client 10.10.1 fails to parse AtCoder's "MiB"
            # memory limit unit — patch it to accept MiB/GiB/KiB as well.
            patchedPython = pkgs.python3.override {
              packageOverrides = self: super: {
                online-judge-api-client = super.online-judge-api-client.overrideAttrs (old: {
                  postPatch =
                    (old.postPatch or "")
                    + ''
                      sed -i "s/(KB|MB)'/(KB|MB|MiB|GiB|KiB)'/g" onlinejudge/service/atcoder.py
                      sed -i "s/elif memory_limit_unit == 'MB':/elif memory_limit_unit in ('MB', 'MiB', 'GiB', 'KiB'):/" onlinejudge/service/atcoder.py
                    '';
                });
              };
            };
          in
          pkgs.mkShell {
          packages = [
            pkgs.go_1_25
            pkgs.gopls
            pkgs.go-task
            pkgs.nodejs
            (patchedPython.withPackages (ps: [
              ps.online-judge-tools
            ]))
          ];

          shellHook = ''
            export NPM_CONFIG_PREFIX="$HOME/.npm-global"
            export PATH="$HOME/.npm-global/bin:$PATH"
            if ! command -v acc &>/dev/null; then
              echo "Installing atcoder-cli..."
              npm install -g atcoder-cli
            fi
            acc config oj-path "$(which oj)"
          '';
        };
      });
    };
}
