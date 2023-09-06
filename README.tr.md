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

<h2 align="center">hyper-v-api-ps</h2>
<b>hyper-v-api-ps</b> Hyper-V sanal makine bilgilerine RESTful API'lar aracılığıyla erişim sağlamak için bir araçtır.
</div>

---

## İçindekiler 

- [İçindekiler](#i̇çindekiler)
- [Kurulum](#kurulum)
- [Kullanım](#kullanım)
- [Lisans](#lisans)

---

## Kurulum

1. `%PROGRAMFILES%\hyper-v-rest-ps` dizinini oluşturun.
2. `hyper-v-rest-ps.exe` dosyasını `%PROGRAMFILES%\hyper-v-rest-ps` dizinine kopyalayın.
3. Yönetici olarak Windows PowerShell'i açın ve aşağıdaki komutları çalıştırın:

```powershell
PS C:\WINDOWS\system32> cd "$env:PROGRAMFILES\hyper-v-rest-ps"
PS C:\Program Files\hyper-v-rest-ps> .\hyper-v-rest-ps.exe --service=install
PS C:\Program Files\hyper-v-rest-ps> .\hyper-v-rest-ps.exe --service=start
```

## Kullanım

- Tüm sanal makineler: `/vms`
- CPU bilgileri: `/vms/<VM-ID>/processor`
- Tüm sanal makinelerin CPU bilgileri: `/vms/all/processor`
- RAM bilgileri: `/vms/<VM-ID>/memory`
- Tüm sanal makinelerin RAM bilgileri: `/vms/all/memory`
- Disk bilgileri: `/vms/<VM-ID>/vhd`
- Tüm sanal makinelerin disk bilgileri: `/vms/all/vhd`
- Ağ bilgileri: `/vms/<VM-ID>/network`
- Tüm sanal makinelerin ağ bilgileri: `/vms/all/network`
- Sürüm: `/version`

---

## Lisans 

hyper-v-rest-ps, GPL-3.0 lisanslıdır. Detaylar için [LICENSE](LICENSE) dosyasına bakınız.

[contributors-shield]: https://img.shields.io/github/contributors/monobilisim/hyper-v-rest-ps.svg?style=for-the-badge
[contributors-url]: https://github.com/monobilisim/hyper-v-rest-ps/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/monobilisim/hyper-v-rest-ps.svg?style=for-the-badge
[forks-url]: https://github.com/monobilisim/hyper-v-rest-ps/network/members
[stars-shield]: https://img.shields.io/github/stars/monobilisim/hyper-v-rest-ps.svg?style=for-the-badge
[stars-url]: https://github.com/monobilisim/hyper-v-rest-ps/stargazers
[issues-shield]: https://img.shields.io/github/issues/monobilisim/hyper-v-rest-ps.svg?style=for-the-badge
[issues-url]: https://github.com/monobilisim/hyper-v-rest-ps/issues
[license-shield]: https://img.shields.io/github/license/monobilisim/hyper-v-rest-ps.svg?style=for-the-badge
[license-url]: https://github.com/monobilisim/hyper-v-rest-ps/blob/master/LICENSE
