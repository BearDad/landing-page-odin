import { defineConfig } from 'vite'

export default defineConfig({
	root: './public',           // Tell Vite that public is your source root
	build: {
		outDir: '../dist',        // Output build to ../dist
		emptyOutDir: true
	},
	server: {
		port: 8888,
		open: true
	}
})
