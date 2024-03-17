package main

import "strings"

func Format(input string) (output string) {
	output = input
	output = strings.ReplaceAll(output, "§0", "\033[0;0m\033[0;30m") // black
	output = strings.ReplaceAll(output, "§1", "\033[0;0m\033[0;34m") // dark_blue
	output = strings.ReplaceAll(output, "§2", "\033[0;0m\033[0;32m") // dark_green
	output = strings.ReplaceAll(output, "§3", "\033[0;0m\033[0;36m") // dark_aqua
	output = strings.ReplaceAll(output, "§4", "\033[0;0m\033[0;31m") // dark_red
	output = strings.ReplaceAll(output, "§5", "\033[0;0m\033[0;35m") // dark_purple
	output = strings.ReplaceAll(output, "§6", "\033[0;0m\033[0;33m") // gold
	output = strings.ReplaceAll(output, "§7", "\033[0;0m\033[0;37m") // gray
	output = strings.ReplaceAll(output, "§8", "\033[0;0m\033[0;90m") // dark_gray
	output = strings.ReplaceAll(output, "§9", "\033[0;0m\033[0;94m") // blue
	output = strings.ReplaceAll(output, "§a", "\033[0;0m\033[0;92m") // green
	output = strings.ReplaceAll(output, "§b", "\033[0;0m\033[0;96m") // aqua
	output = strings.ReplaceAll(output, "§c", "\033[0;0m\033[0;91m") // red
	output = strings.ReplaceAll(output, "§d", "\033[0;0m\033[0;95m") // light_purple
	output = strings.ReplaceAll(output, "§e", "\033[0;0m\033[0;93m") // yellow
	output = strings.ReplaceAll(output, "§f", "\033[0;0m\033[0;97m") // white
	output = strings.ReplaceAll(output, "§k", "\033[0;0m\033[0;97m") // obfuscated/MTS
	output = strings.ReplaceAll(output, "§l", "\033[0;0m\033[0;97m") // bold
	output = strings.ReplaceAll(output, "§m", "\033[0;0m\033[0;97m") // strikethrough
	output = strings.ReplaceAll(output, "§n", "\033[0;0m\033[0;97m") // underline
	output = strings.ReplaceAll(output, "§o", "\033[0;0m\033[0;97m") // italic
	output = strings.ReplaceAll(output, "§r", "\033[0;0m\033[0;97m") // rese
	return
}