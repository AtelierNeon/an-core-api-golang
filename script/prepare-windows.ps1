##
## Built-in config
##
$InformationPreference = 'Continue'
#$VerbosePreference = 'Continue'

##
## Global config
##
$ChocoCli = 'choco'
$GoCli = 'go'

##
## My variables
##
$MySystemChocoCliFound = $false
$MySystemGoCliFound = $false



## Detect Chocolatey
$MyChocoProcess = $null
$MyChocoProcessHandle = $null
if (-not $MySystemChocoCliFound) {
    Write-Information "[PowerShell] Detecting Chocolatey ..."
    try {
        $MyChocoProcess = Start-Process -FilePath "${Env:ProgramData}\chocolatey\$ChocoCli" -WindowStyle Hidden -PassThru `
            -ArgumentList "--version"
        $MyChocoProcessHandle = $MyChocoProcess.Handle
        $MyChocoProcess.WaitForExit()
        $MyChocoProcessExitCode = $MyChocoProcess.ExitCode
        if ($MyChocoProcessExitCode -ne 0) {
            Write-Information "[PowerShell] Detecting Chocolatey ... INCORRECT (ExitCode: $MyChocoProcessExitCode)"
            $MySystemChocoCliFound = $false
        }
        else {
            Write-Information "[PowerShell] Detecting Chocolatey ... FOUND"
            $MySystemChocoCliFound = $true
        }
    }
    catch {
        Write-Information "[PowerShell] Detecting Chocolatey ... NOT FOUND"
        $MySystemChocoCliFound = $false
    }
    finally {
        if ($null -ne $MyChocoProcessHandle) {
            $MyChocoProcessHandle = $null
        }
        if ($null -ne $MyChocoProcess) {
            $MyChocoProcess.Dispose()
            $MyChocoProcess = $null
        }
    }
}



## Install Chocolatey
if (-not $MySystemChocoCliFound) {
    Write-Information "[PowerShell] Installing Chocolatey ..."
    try {
        [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072
        Invoke-Expression ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
    }
    catch {
        Write-Error "[PowerShell] Installing Chocolatey ... FAILED"
        Exit 1
    }
    Write-Information "[PowerShell] Installing Chocolatey ... DONE"
}



## Append Chocolatey shim folder into PATH
$Env:PATH = $Env:PATH + ';' + $Env:ProgramData + '\chocolatey\bin'



## Detect GO SDK
$MyGoProcess = $null
$MyGoProcessHandle = $null
if (-not $MySystemGoCliFound) {
    Write-Information "[PowerShell] Detecting GO SDK ..."
    try {
        $MyGoProcess = Start-Process -FilePath "$GoCli" -WindowStyle Hidden -PassThru `
            -ArgumentList "version"
        $MyGoProcessHandle = $MyGoProcess.Handle
        $MyGoProcess.WaitForExit()
        $MyGoProcessExitCode = $MyGoProcess.ExitCode
        if ($MyGoProcessExitCode -ne 0) {
            Write-Information "[PowerShell] Detecting GO SDK ... NOT FOUND (ExitCode: $MyGoProcessExitCode)"
            $MySystemGoCliFound = $false
        }
        else {
            Write-Information "[PowerShell] Detecting GO SDK ... FOUND"
            $MySystemGoCliFound = $true
        }
    }
    catch {
        Write-Information "[PowerShell] Detecting GO SDK ... NOT FOUND"
        $MySystemGoCliFound = $false
    }
    finally {
        if ($null -ne $MyGoProcessHandle) {
            $MyGoProcessHandle = $null
        }
        if ($null -ne $MyGoProcess) {
            $MyGoProcess.Dispose()
            $MyGoProcess = $null
        }
    }
}



## Install GO SDK
if (-not $MySystemGoCliFound) {
    Write-Information "[PowerShell] Installing GO SDK ..."
    try {
        $MyChocoProcess = Start-Process -FilePath "${Env:ProgramData}\chocolatey\$ChocoCli" -WindowStyle Hidden -PassThru `
            -ArgumentList "install golang --force -y"
        $MyChocoProcessHandle = $MyChocoProcess.Handle
        $MyChocoProcess.WaitForExit()
        $MyChocoProcessExitCode = $MyChocoProcess.ExitCode
        if ($MyChocoProcessExitCode -ne 0) {
            Write-Error "[PowerShell] Installing GO SDK ... INCORRECT (ExitCode: $MyChocoProcessExitCode)"
            Exit 1
        }
    }
    catch {
        Write-Error "[PowerShell] Installing GO SDK ... FAILED"
        Exit 1
    }
    finally {
        if ($null -ne $MyChocoProcessHandle) {
            $MyChocoProcessHandle = $null
        }
        if ($null -ne $MyChocoProcess) {
            $MyChocoProcess.Dispose()
            $MyChocoProcess = $null
        }
    }
    Write-Information "[PowerShell] Installing GO SDK ... DONE"
}
