aws-assume-role
===

[![Build Status](https://travis-ci.org/mitchelldavis/aws-assume-role.svg?branch=master)](https://travis-ci.org/mitchelldavis/aws-assume-role)

This tool can be used to assume a role in an AWS account.  I needed this tool because I would like to sepperate out Production and Development environments into separate AWS accounts.  If my continuous delivery solution is in my Development account, I need to be able to assume a role in the Production account in order to deploy my software.

There are tools out there that export AWS credentials to the current shell, but I wanted a single executable. Python and shell scripts are hard to deal with in an environment where you are downloading your dependencies.

Use
---

The tool is very simple.  I literally copied and pasted the [example](https://docs.aws.amazon.com/sdk-for-go/api/service/sts/#STS.AssumeRole) found in the [AWS Golang SDK]().  Of course, I added in some flag parsing, error handling, and the output to the consule but that example is pretty much it.

```
Usage of ./.bin/aws-assume-role:
  -duration int
        The amount of time, in seconds, to keep the temporary credentials alive. (default 3600)
  -externalId string
        A unique identifier that may be required to assume the role.
  -roleArn string
        The Arn of the role to assume.
  -sessionName string
        A unique identifier for the session so you can distinguish between to prinipals assuming the same role.
```

Build
---

```
git clone git:github.com/mitchelldavis/aws-assume-role
cd aws-assume-role
make
```

LICENSE
---

This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <https://unlicense.org>
