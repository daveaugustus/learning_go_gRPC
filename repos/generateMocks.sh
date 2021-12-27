#!/bin/bash

echo "Generating mocks for repos"
echo ""

REPOS_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

ALL_REPOS=(
    "UsersRepo"
    "GlobalRepository"
)

for repo in "${ALL_REPOS[@]}"
do
    echo "Generating mock for $repo"
    dest="$REPOS_ROOT/mocks/$repo.go"
    src="github.com/davetweetlive/learning_go_gRPC/repos"
    GO111MODULE=on mockgen -destination $dest $src $repo
done

echo ""
echo "Finished generating mocks!"
exit 0
