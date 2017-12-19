package commands

func Unzip(zipFile, destDir string) {
	run("unzip", "-qq", zipFile, "-d", destDir)
}
