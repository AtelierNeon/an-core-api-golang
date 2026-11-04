##
## Built-in config
##
$InformationPreference = 'Continue'
#$VerbosePreference = 'Continue'

##
## Global config
##
$GoCli = 'go'

##
## Project config
##
####
#### Project level config
####
$ProjectGoVersion = if ($Env:MY_PROJECT_GO_VERSION) { $Env:MY_PROJECT_GO_VERSION } else { '1.23.12' }
$ProjectUseDefaultGoCli = if ($Env:MY_PROJECT_USE_DEFAULT_GO_CLI) { $Env:MY_PROJECT_USE_DEFAULT_GO_CLI } else { 'OFF' }

##
## My Variables
##
$MyProjectGoCli = $GoCli + $ProjectGoVersion
$MyProjectGoCliDetected = $false
$MyProjectGoVersion = $ProjectGoVersion
if ('ON'.Equals($ProjectUseDefaultGoCli)) {
    $MyProjectGoCli = $GoCli
    $MyProjectGoCliDetected = $true
    $MyProjectGoVersion = 'default'
}



## Print build information
Write-Information "[PowerShell] Project information: GO version: `"$MyProjectGoVersion`""



## Detect GO SDK
$MyGoProcess = $null
$MyGoProcessHandle = $null
Write-Information "[PowerShell] Detecting GO SDK ..."
try {
    $MyGoProcess = Start-Process -FilePath "$GoCli" -WindowStyle Hidden -PassThru `
        -ArgumentList "version"
    $MyGoProcessHandle = $MyGoProcess.Handle
    $MyGoProcess.WaitForExit()
    $MyGoProcessExitCode = $MyGoProcess.ExitCode
    if ($MyGoProcessExitCode -ne 0) {
        Write-Error "[PowerShell] Detecting GO SDK ... NOT FOUND (ExitCode: $MyGoProcessExitCode)"
        Exit 1
    }
}
catch {
    Write-Error "[PowerShell] Detecting GO SDK ... NOT FOUND"
    Exit 1
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
Write-Information "[PowerShell] Detecting GO SDK ... FOUND"



## Detect GOPATH
$MyGoPath = ''
Write-Information "[PowerShell] Detecting GOPATH ..."
try {
    $MyGoProcessStartInfo = New-Object System.Diagnostics.ProcessStartInfo
    $MyGoProcessStartInfo.FileName = $GoCli
    $MyGoProcessStartInfo.RedirectStandardError = $true
    $MyGoProcessStartInfo.RedirectStandardOutput = $true
    $MyGoProcessStartInfo.UseShellExecute = $false
    $MyGoProcessStartInfo.Arguments = "env GOPATH"
    $MyGoProcess = New-Object System.Diagnostics.Process
    $MyGoProcess.StartInfo = $MyGoProcessStartInfo
    $MyGoProcess.Start() | Out-Null
    $MyGoProcessHandle = $MyGoProcess.Handle
    $MyGoProcess.WaitForExit()
    $MyGoProcessExitCode = $MyGoProcess.ExitCode
    $MyGoPath = $MyGoProcess.StandardOutput.ReadToEnd()
    if ($MyGoProcessExitCode -ne 0) {
        Write-Error "[PowerShell] Detecting GOPATH ... FAILED (ExitCode: $MyGoProcessExitCode)"
        Exit 1
    }
    $MyGoPath = $MyGoPath.Trim()
    if (''.Equals($MyGoPath)) {
        Write-Error "[PowerShell] Detecting GOPATH ... FAILED (Empty)"
        Exit 1
    }
}
catch {
    Write-Error "[PowerShell] Detecting GOPATH ... FAILED"
    Exit 1
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
Write-Information "[PowerShell] Detecting GOPATH ... DONE"



## Add GOPATH\bin into PATH
Write-Information "[PowerShell] Adding GOPATH\bin into PATH ..."
$MyGoPathBin = "$MyGoPath\bin"
if ($Env:Path -contains $MyGoPathBin) {
    Write-Information "[PowerShell] Adding GOPATH\bin into PATH ... SKIPPED"
}
else {
    $Env:Path = "$MyGoPathBin;" + $Env:Path
    Write-Information "[PowerShell] Adding GOPATH\bin into PATH ... DONE"
}



## Detect project GO SDK
Write-Information "[PowerShell] Detecting project GO SDK ..."
if (-not $MyProjectGoCliDetected) {
    $MyGoProcess = $null
    $MyGoProcessHandle = $null
    try {
        $MyGoProcess = Start-Process -FilePath "$MyProjectGoCli" -WindowStyle Hidden -PassThru `
            -ArgumentList "version"
        $MyGoProcessHandle = $MyGoProcess.Handle
        $MyGoProcess.WaitForExit()
        $MyGoProcessExitCode = $MyGoProcess.ExitCode
        if ($MyGoProcessExitCode -ne 0) {
            Write-Information "[PowerShell] Detecting project GO SDK ... NOT FOUND (ExitCode: $MyGoProcessExitCode)"
            $MyProjectGoCliDetected = $false
        }
        else {
            Write-Information "[PowerShell] Detecting project GO SDK ... FOUND"
            $MyProjectGoCliDetected = $true
        }
    }
    catch {
        Write-Information "[PowerShell] Detecting project GO SDK ... NOT FOUND (project GO SDK is missing)"
        $MyProjectGoCliDetected = $false
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



## Restore project GO SDK
if (-not $MyProjectGoCliDetected) {
    Write-Information "[PowerShell] Restoring project GO SDK ..."

    ## Restore project GO SDK - Create SDK downloader
    $MyGoProcess = $null
    $MyGoProcessHandle = $null
    Write-Information "[PowerShell] Restoring project GO SDK ... Creating SDK downloader ..."
    try {
        $MyGoProcess = Start-Process -FilePath "$GoCli" -WindowStyle Hidden -PassThru `
            -ArgumentList "install golang.org/dl/go$ProjectGoVersion@latest"
        $MyGoProcessHandle = $MyGoProcess.Handle
        $MyGoProcess.WaitForExit()
        $MyGoProcessExitCode = $MyGoProcess.ExitCode
        if ($MyGoProcessExitCode -ne 0) {
            Write-Error "[PowerShell] Restoring project GO SDK ... Creating SDK downloader ... FAILED (ExitCode: $MyGoProcessExitCode)"
            Exit 1
        }
    }
    catch {
        Write-Error "[PowerShell] Restoring project GO SDK ... Creating SDK downloader ... FAILED (Default GO SDK is missing)"
        Exit 1
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
    Write-Information "[PowerShell] Restoring project GO SDK ... Creating SDK downloader ... DONE"

    ## Restore project GO SDK - Download SDK
    $MyGoProcess = $null
    $MyGoProcessHandle = $null
    Write-Information "[PowerShell] Restoring project GO SDK ... Downloading SDK ..."
    try {
        $MyGoProcess = Start-Process -FilePath "$MyProjectGoCli" -WindowStyle Hidden -PassThru `
            -ArgumentList "download"
        $MyGoProcessHandle = $MyGoProcess.Handle
        $MyGoProcess.WaitForExit()
        $MyGoProcessExitCode = $MyGoProcess.ExitCode
        if ($MyGoProcessExitCode -ne 0) {
            Write-Error "[PowerShell] Restoring project GO SDK ... Downloading SDK ... FAILED (ExitCode: $MyGoProcessExitCode)"
            Exit 1
        }
    }
    catch {
        Write-Error "[PowerShell] Restoring project GO SDK ... Downloading SDK ... FAILED (GO SDK downloader is missing)"
        Exit 1
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
    Write-Information "[PowerShell] Restoring project GO SDK ... Downloading SDK ... DONE"
}



## Build project
Write-Information "[PowerShell] Building project ..."

## Build project - Fetch project dependencies
$MyGoProcess = $null
$MyGoProcessHandle = $null
Write-Information "[PowerShell] Building project ... Fetching project dependencies ..."
try {
    $MyGoProcess = Start-Process -FilePath "$MyProjectGoCli" -NoNewWindow -PassThru `
        -ArgumentList "mod tidy"
    $MyGoProcessHandle = $MyGoProcess.Handle
    $MyGoProcess.WaitForExit()
    $MyGoProcessExitCode = $MyGoProcess.ExitCode
    if ($MyGoProcessExitCode -ne 0) {
        Write-Error "[PowerShell] Building project ... Fetching project dependencies ... FAILED (ExitCode: $MyGoProcessExitCode)"
        Exit 1
    }
}
catch {
    Write-Error "[PowerShell] Building project ... Fetching project dependencies ... FAILED (Default GO SDK is missing)"
    Exit 1
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
Write-Information "[PowerShell] Building project ... Fetching project dependencies ... DONE"

## Build project - Run go test
Write-Information "[PowerShell] Building Project ... Running go test ..."
$Env:GOARCH = ''
$MyGoProcess = $null
$MyGoProcessHandle = $null
try {
    $MyGoProcess = Start-Process -FilePath "$MyProjectGoCli" -NoNewWindow -PassThru `
        -ArgumentList "test -short -cover ./..."
    $MyGoProcessHandle = $MyGoProcess.Handle
    $MyGoProcess.WaitForExit()
    $MyGoProcessExitCode = $MyGoProcess.ExitCode
    if ($MyGoProcessExitCode -ne 0) {
        Write-Error "[PowerShell] Building Project ... Running go test ... FAILED (ExitCode: $MyGoProcessExitCode)"
        Exit 1
    }
}
catch {
    Write-Error "[PowerShell] Building Project ... Running go test ... FAILED (Project GO SDK is missing)"
    Exit 1
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
Write-Information "[PowerShell] Building Project ... Running go test ... DONE"

Write-Information "[PowerShell] Building project ... DONE"
