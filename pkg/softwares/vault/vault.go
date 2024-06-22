package vault

type VaultInstaller struct{}

// func (vi *VaultInstaller) Install(name, version, arch, platform string, notIf ...bool) {
// func (vi *VaultInstaller) Install(name, archiveURL, archiveFilename, destFolder string, notIf ...bool) {
// 	if notIf[0] {
// 		fmt.Printf("[%v][Install] Ignore Install due to NotIf\n", name)
// 	} else {
// 		fmt.Printf("[%v][Install] Download archive\n", name)

// 		err := file.Download(
// 			archiveURL,
// 			archiveFilename,
// 		)
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Printf("[%v][Install] Decompress archive\n", name)
// 		err = archive.Unzip(archiveFilename, destFolder)
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Printf("[%v][Install] Delete archive\n", name)
// 		err = file.Delete(archiveFilename)
// 		if err != nil {
// 			panic(err)
// 		}

// 		// Rename file to vault_${version}
// 		err = os.Rename(
// 			fmt.Sprintf(
// 				"%v/vault",
// 				destFolder,
// 			),
// 			fmt.Sprintf(
// 				"%v/vault_%v",
// 				destFolder,
// 				version,
// 			),
// 		)
// 		if err != nil {
// 			panic(err)
// 		}

// 		// Link vault_${version} to vault
// 		err = os.Symlink(
// 			fmt.Sprintf(
// 				"%v/vault_%v",
// 				destFolder,
// 				version,
// 			),
// 			fmt.Sprintf(
// 				"%v/vault",
// 				destFolder,
// 			),
// 		)
// 		if err != nil {
// 			panic(err)
// 		}

// 		// Delete LICENSE.txt file (from archive)
// 		err = file.Delete(fmt.Sprintf("%v/LICENSE.txt", destFolder))
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Printf("[%v][Install] Complete\n", name)
// 	}
// }
