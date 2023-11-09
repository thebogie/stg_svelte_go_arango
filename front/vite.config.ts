import { purgeCss } from 'vite-plugin-tailwind-purgecss';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	server: {
		port: 50003,
		host: '0.0.0.0',

	},
	plugins: [sveltekit(), purgeCss()]
});
