## Kurulum

* `%PROGRAMFILES%\wmi-rest` klasörü oluşturulur.
* `wmi-rest.exe` dosyası, `%PROGRAMFILES%\wmi-rest` klasörüne kopyalanır.
* `config.yml` dosyası oluşturulur ve `%PROGRAMFILES%\wmi-rest` klasörüne kaydedilir. (Bkz. [config.sample.yml](conf/config.sample.yml))
* Windows PowerShell, yönetici olarak açılır ve aşağıdaki komutlar çalıştırılır:

       PS C:\WINDOWS\system32> cd "$env:PROGRAMFILES\wmi-rest"
       PS C:\Program Files\wmi-rest> .\wmi-rest.exe --service=install
       PS C:\Program Files\wmi-rest> .\wmi-rest.exe --service=start

## Kullanım

* VM'leri görüntüleme: `/vms`
  * VM bilgileri, [`MSVM_ComputerSystem`](https://learn.microsoft.com/en-us/windows/win32/hyperv_v2/msvm-computersystem) sınıfının özelliklerinden geliyor.
* Tek bir VM'i görüntüleme: `/vms/<VM adı>`
  * VM adı, [`MSVM_ComputerSystem`](https://learn.microsoft.com/en-us/windows/win32/hyperv_v2/msvm-computersystem) nesnesinin `Name` özelliğinde bulunuyor.
* VM'in memory bilgisini görüntüleme `/vms/<VM adı>/memory`
  * Memory bilgileri, [`MSVM_Memory`](https://learn.microsoft.com/en-us/previous-versions/windows/desktop/virtual/msvm-memory) sınıfının özelliklerinden geliyor.
* VM'in processor bilgisini görüntüleme `/vms/<VM adı>/processor`
  * Processor bilgileri, [`MSVM_Processor`](https://learn.microsoft.com/en-us/windows/win32/hyperv_v2/msvm-processor) sınıfının özelliklerinden geliyor.
* VM'in storage bilgisini görüntüleme `/vms/<VM adı>/processor`
  * Storage bilgileri, [`MSVM_DiskDrive`](https://learn.microsoft.com/en-us/windows/win32/hyperv_v2/msvm-diskdrive) sınıfının özelliklerinden geliyor.
