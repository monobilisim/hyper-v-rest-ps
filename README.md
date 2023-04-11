## Kurulum

- `%PROGRAMFILES%\wmi-rest` klasörü oluşturulur.
- `wmi-rest.exe` dosyası, `%PROGRAMFILES%\wmi-rest` klasörüne kopyalanır.
- Windows PowerShell, yönetici olarak açılır ve aşağıdaki komutlar çalıştırılır:

       PS C:\WINDOWS\system32> cd "$env:PROGRAMFILES\wmi-rest"
       PS C:\Program Files\wmi-rest> .\wmi-rest.exe --service=install
       PS C:\Program Files\wmi-rest> .\wmi-rest.exe --service=start

## Kullanım

- VM'leri görüntüleme: `/vms`
- Tek bir VM'i görüntüleme: `/vms/<VM adı>`
- VM'in storage bilgisini görüntüleme `/vms/<VM adı>/vhd`
- Wmi-rest versiyonunu görüntüleme: `/version`
