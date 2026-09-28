```yml
Run actions/github-script@v9
  with:
    script: const { run } = require('./.github/scripts/update-pr-template.go');
  return run({ github, context, core });
  
    github-token: ***
    debug: false
    user-agent: actions/github-script
    result-encoding: json
    retries: 0
    retry-exempt-status-codes: 400,401,403,404,422
/home/runner/work/custom-repo-template/custom-repo-template/.github/scripts/update-pr-template.go:9
package main
        ^^^^
SyntaxError: Unexpected identifier 'main'
    at wrapSafe (node:internal/modules/cjs/loader:1804:18)
    at Module._compile (node:internal/modules/cjs/loader:1846:20)
    at Object..js (node:internal/modules/cjs/loader:2003:10)
    at Module.load (node:internal/modules/cjs/loader:1594:32)
    at Module._load (node:internal/modules/cjs/loader:1396:12)
    at wrapModuleLoad (node:internal/modules/cjs/loader:255:19)
    at Module.require (node:internal/modules/cjs/loader:1617:12)
    at require (node:internal/modules/helpers:153:16)
    at Object.apply (/home/runner/work/_actions/actions/github-script/v9/dist/index.js:65038:27)
    at eval (eval at callAsyncFunction (/home/runner/work/_actions/actions/github-script/v9/dist/index.js:64949:16), <anonymous>:3:17)
Error: Unhandled error: SyntaxError: Unexpected identifier 'main'
```