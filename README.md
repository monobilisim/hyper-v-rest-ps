[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![GPL License][license-shield]][license-url]

[![Readme in English](https://img.shields.io/badge/Readme-English-blue)](README.md)
[![Readme in Turkish](https://img.shields.io/badge/Readme-Turkish-red)](README.tr.md)

<div align="center"> 
<a href="https://mono.net.tr/">
  <img src="https://monobilisim.com.tr/images/mono-bilisim.svg" width="340"/>
</a>

<h2 align="center">wmi-rest</h2>
<b>wmi-rest</b> is a tool for accessing Hyper-V VM information through RESTful APIs.
</div>

---

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Installation](#installation)
- [Usage](#usage)
- [License](#license)

---

## Installation

1. Create the `%PROGRAMFILES%\wmi-rest` directory.
2. Copy the `wmi-rest.exe` file to the `%PROGRAMFILES%\wmi-rest` directory.
3. Open Windows PowerShell as an administrator and execute the following commands:

```powershell
PS C:\WINDOWS\system32> cd "$env:PROGRAMFILES\wmi-rest"
PS C:\Program Files\wmi-rest> .\wmi-rest.exe --service=install
PS C:\Program Files\wmi-rest> .\wmi-rest.exe --service=start
```

## Usage

- All virtual machines: `/vms`
- CPU information: `/vms/<VM-ID>/processor`
- CPU information for all virtual machines: `/vms/all/processor`
- RAM information: `/vms/<VM-ID>/memory`
- RAM information for all virtual machines: `/vms/all/memory`
- Disk information: `/vms/<VM-ID>/vhd`
- Disk information for all virtual machines: `/vms/all/vhd`
- Network information: `/vms/<VM-ID>/network`
- Network information for all virtual machines: `/vms/all/network`
- Version: `/version`

---

## License

wmi-rest is GPL-3.0 licensed. See [LICENSE](LICENSE) file for details.

[contributors-shield]: https://img.shields.io/github/contributors/monobilisim/wmi-rest.svg?style=for-the-badge
[contributors-url]: https://github.com/monobilisim/wmi-rest/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/monobilisim/wmi-rest.svg?style=for-the-badge
[forks-url]: https://github.com/monobilisim/wmi-rest/network/members
[stars-shield]: https://img.shields.io/github/stars/monobilisim/wmi-rest.svg?style=for-the-badge
[stars-url]: https://github.com/monobilisim/wmi-rest/stargazers
[issues-shield]: https://img.shields.io/github/issues/monobilisim/wmi-rest.svg?style=for-the-badge
[issues-url]: https://github.com/monobilisim/wmi-rest/issues
[license-shield]: https://img.shields.io/github/license/monobilisim/wmi-rest.svg?style=for-the-badge
[license-url]: https://github.com/monobilisim/wmi-rest/blob/master/LICENSE