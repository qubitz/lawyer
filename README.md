# Lawyer ⚖️
![Chocolatey Version](https://img.shields.io/chocolatey/v/lawyer) [![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/qubitz/lawyer/LICENSE.txt)

A tool for maintaining file headers in egregious violation of strict standards.

## Pro bono services

### Indictments 🔍

Analyze files to find headers that don't meet the strict criteria as specified in
the [Law](#terms).

```
lawyer indict file.txt # single file
lawyer indict *.txt    # or multiple files with globbing
```
Output:
```
searching for the law...found ./law.yaml
fail.txt is guilty

evidence
--------------------------------------------------
/// <copyright file="file.txt" company="ABC">
/// ABC © 2021. All rights reserved.
/// </copyright
--------------------------------------------------

indictment complete
```

## Procurement process

### [Chocolatey](https://community.chocolatey.org/packages/lawyer)

like yum or apt-get, but for Windows
```
choco install lawyer
``` 

## Terms

Term | Definition
---- | ----------
Law | A file that specifies how file headers should look. See `lawyer help law` for more information.