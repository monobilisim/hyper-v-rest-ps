## Kurulum

- `%PROGRAMFILES%\wmi-rest` klasörü oluşturulur.
- `wmi-rest.exe` dosyası, `%PROGRAMFILES%\wmi-rest` klasörüne kopyalanır.
- Windows PowerShell, yönetici olarak açılır ve aşağıdaki komutlar çalıştırılır:

       PS C:\WINDOWS\system32> cd "$env:PROGRAMFILES\wmi-rest"
       PS C:\Program Files\wmi-rest> .\wmi-rest.exe --service=install
       PS C:\Program Files\wmi-rest> .\wmi-rest.exe --service=start

## Kullanım

- Tüm VM'ler: `/vms`
- CPU ve GuestOS: `/vms/<VM>/processor`
- Memory: `/vms/<VM>/memory`
- Disk bilgisini `/vms/<VM>/vhd`
- Versiyon: `/version`
