const path = require("path")
const fs = require("fs")
const { exec } = require("child_process")


const branch = new Promise((resolve, reject) => exec("git branch --show-current", (err, stdout, stderr) => {
	if (stdout) resolve(stdout.trim())
	else if (stderr) reject(stderr)
	else reject(new Error("git branch not found"))
})
)
// vercel add env from .env.local

const pull = async () => {
	const file = fs.readFileSync(path.resolve(__dirname, "../.env.local"), "utf-8")
	const lines = file.split("\n").map(line => line.trim()).filter(line => {
		if (line.length < 0) return false
		if (line.startsWith("#")) return false
		if (!line.includes("=")) return false
		if (line.includes("VERCEL")) return false
		if (line.includes("#")) line = line.split("#")[0]
		return true
	})
	.map(async line => {
		const [name, value] = line.split("=")
		// get git branch
		// run vercel env add ${name} ${value}
		exec(`vercel env add ${name} "${value}" ${await branch}`, (err, stdout, stderr) => {
			if (err) {
				console.error(err)
				// stderr && console.error(stderr)
				return
			}
			console.log("env added", "[name]", name, "[value]", value)
			// console.log(stdout)
		})
	})

	return lines
}

pull().then(
	() => {
		console.log("All environment variables successfully pushed to vercel 🎆")
		process.exit(0)
	}
)



