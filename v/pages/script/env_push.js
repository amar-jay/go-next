const path = require("path")
const fs = require("fs")
const { exec } = require("child_process")


// vercel add env from .env.local

const pull = async () => {
	const file = fs.readFileSync(path.resolve(__dirname, "../.env.local"), "utf-8")
	const lines = file.split("\n").map(line => line.trim()).filter(line => line.length > 0)
	.map(line => {
		const [name, value] = line.split("=")
		// get git branch
		const branch = exec("git branch --show-current", (err, stdout, stderr) => {
			return stdout
		})
		if (!branch) {
			console.error("git branch not found")
			return
		}
		// run vercel env add ${name} ${value}
		exec(`vercel env add ${name} ${value} ${branch}`, (err, stdout, stderr) => {
			if (err) {
				console.error(err)
				return
			}
			console.log("env added", "[name]", name, "[value]", value)
			// console.log(stdout)
		})
	})

	return lines
}

pull()


