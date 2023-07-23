## Kurulum

- `%PROGRAMFILES%\wmi-rest` klasörü oluşturulur.
- `wmi-rest.exe` dosyası, `%PROGRAMFILES%\wmi-rest` klasörüne kopyalanır.
- Windows PowerShell, yönetici olarak açılır ve aşağıdaki komutlar çalıştırılır:

       PS C:\WINDOWS\system32> cd "$env:PROGRAMFILES\wmi-rest"
       PS C:\Program Files\wmi-rest> .\wmi-rest.exe --service=install
       PS C:\Program Files\wmi-rest> .\wmi-rest.exe --service=start

## Kullanım

- Tüm sanal makineler: `/vms`
- CPU bilgisi: `/vms/<VM-ID>/processor`
- Tüm sanal makinelerin CPU bilgisi: `/vms/all/processor`
- RAM bilgisi: `/vms/<VM-ID>/memory`
- Tüm sanal makinelerin RAM bilgisi: `/vms/all/memory`
- Disk bilgisi: `/vms/<VM-ID>/vhd`
- Tüm sanal makinelerin disk bilgisi: `/vms/all/vhd`
- Ağ bilgisi: `/vms/<VM-ID>/network`
- Tüm sanal makinelerin ağ bilgisi: `/vms/all/network`
- Versiyon: `/version`
