{
  description = "Flake for hrpc";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    hrpc.url = "github:harmony-development/hrpc/main";
    flakeUtils.url = "github:numtide/flake-utils";
    flakeCompat = {
      url = "github:edolstra/flake-compat";
      flake = false;
    };
  };

  outputs = inputs:
    with inputs.flakeUtils.lib; eachSystem [ "x86_64-linux" ] (system:
      let
        pkgs = import inputs.nixpkgs {
          inherit system;
        };
        hrpcPkgs = inputs.hrpc.packages.${system};
      in
      {
          devShell =
            pkgs.mkShell {
              name = "inviter-shell";
              buildInputs = with pkgs; [ go buf protobuf hrpcPkgs.protoc-gen-go-hrpc hrpcPkgs.protoc-gen-hrpc ];
            };
      }
    );
}
