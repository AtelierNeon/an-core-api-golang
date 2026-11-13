##
## Built-in config
##
$InformationPreference = 'Continue'
#$VerbosePreference = 'Continue'

##
## Global config
##
$GoCli = 'go'
$ProjectFolder = Join-Path -Path $PSScriptRoot -ChildPath '..'

##
## Project config
##
####
#### Project-level config
####
$ProjectGoVersion = if ($Env:MY_PROJECT_GO_VERSION) { $Env:MY_PROJECT_GO_VERSION } else { '1.23.12' }
$ProjectUseDefaultGoCli = if ($Env:MY_PROJECT_USE_DEFAULT_GO_CLI) { $Env:MY_PROJECT_USE_DEFAULT_GO_CLI } else { 'OFF' }

##
## My variables
##
$MyProjectGoBin = ''
$MyProjectGoBinDetected = $false
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



## Detect GOBIN
if (-not $MyProjectGoBinDetected) {
    Write-Information "[PowerShell] Detecting GOBIN ..."
    try {
        $MyGoProcessStartInfo = New-Object System.Diagnostics.ProcessStartInfo
        $MyGoProcessStartInfo.FileName = $GoCli
        $MyGoProcessStartInfo.RedirectStandardError = $true
        $MyGoProcessStartInfo.RedirectStandardOutput = $true
        $MyGoProcessStartInfo.UseShellExecute = $false
        $MyGoProcessStartInfo.Arguments = "env GOBIN"
        $MyGoProcess = New-Object System.Diagnostics.Process
        $MyGoProcess.StartInfo = $MyGoProcessStartInfo
        $MyGoProcess.Start() | Out-Null
        $MyGoProcessHandle = $MyGoProcess.Handle
        $MyGoProcess.WaitForExit()
        $MyGoProcessExitCode = $MyGoProcess.ExitCode
        $MyGoBin = $MyGoProcess.StandardOutput.ReadToEnd()
        if ($MyGoProcessExitCode -ne 0) {
            Write-Information "[PowerShell] Detecting GOBIN ... INCORRECT (ExitCode: $MyGoProcessExitCode)"
            $MyProjectGoBinDetected = $false
        }
        else {
            $MyGoBin = $MyGoBin.Trim()
            if (''.Equals($MyGoBin)) {
                Write-Information "[PowerShell] Detecting GOBIN ... NOT FOUND"
                $MyProjectGoBinDetected = $false
            }
            else {
                Write-Information "[PowerShell] Detecting GOBIN ... FOUND ( $MyGoBin )"
                $MyProjectGoBin = $MyGoBin
                $MyProjectGoBinDetected = $true
            }
        }
    }
    catch {
        Write-Information "[PowerShell] Detecting GOBIN ... NOT FOUND"
        $MyProjectGoBinDetected = $false
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



## Detect GOPATH
if (-not $MyProjectGoBinDetected) {
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
            Write-Information "[PowerShell] Detecting GOPATH ... INCORRECT (ExitCode: $MyGoProcessExitCode)"
            $MyProjectGoBinDetected = $false
        }
        else {
            $MyGoPath = $MyGoPath.Trim()
            if (''.Equals($MyGoPath)) {
                Write-Information "[PowerShell] Detecting GOPATH ... NOT FOUND"
                $MyProjectGoBinDetected = $false
            }
            else {
                Write-Information "[PowerShell] Detecting GOPATH ... FOUND ( $MyGoPath )"
                $MyProjectGoBin = "$MyGoPath\bin"
                $MyProjectGoBinDetected = $true
            }
        }
    }
    catch {
        Write-Information "[PowerShell] Detecting GOPATH ... NOT FOUND"
        $MyProjectGoBinDetected = $false
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
if (-not $MyProjectGoBinDetected) {
    Write-Error "[PowerShell] Detecting GOBIN or GOPATH ... NOT FOUND"
    Exit 1
}



## Add GOBIN or GOPATH/bin into PATH
Write-Information "[PowerShell] Adding GOBIN or GOPATH/bin into PATH ..."
if ($Env:Path -contains "$MyProjectGoBin") {
    Write-Information "[PowerShell] Adding GOBIN or GOPATH/bin into PATH ... SKIPPED"
}
else {
    $Env:Path = "$MyProjectGoBin;" + $Env:Path
    Write-Information "[PowerShell] Adding GOBIN or GOPATH/bin into PATH ... DONE"
}



## Detect project GO SDK
if (-not $MyProjectGoCliDetected) {
    Write-Information "[PowerShell] Detecting project GO SDK ..."
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
Set-Location -LiteralPath $ProjectFolder -ErrorAction Stop

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
Write-Information "[PowerShell] Building project ... Running go test ..."
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
        Write-Error "[PowerShell] Building project ... Running go test ... FAILED (ExitCode: $MyGoProcessExitCode)"
        Exit 1
    }
}
catch {
    Write-Error "[PowerShell] Building project ... Running go test ... FAILED (Project GO SDK is missing)"
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
Write-Information "[PowerShell] Building project ... Running go test ... DONE"

Write-Information "[PowerShell] Building project ... DONE"
