import { purgeCss } from 'vite-plugin-tailwind-purgecss';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	resolve: {
		alias: {
			'@': '/src', // Alias for the root of the 'src' directory
			'@components': '/src/components', // Alias for the 'components' directory
			'@lib': '/src/lib', // Alias for the 'lib' directory
			'@assets': '/src/assets', // Alias for the 'assets' directory
			'@routes': '/src/routes', // Alias for the 'routes' directory
		},
	},
	server: {
		port: 50003,
		host: '0.0.0.0',

	},
	plugins: [sveltekit(), purgeCss()]
});
