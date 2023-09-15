import { sveltekit } from '@sveltejs/kit/vite';
import type { UserConfig } from 'vite';

const config: UserConfig = {
    server: {
        port: 50003,
        host: '0.0.0.0',

    },
    plugins: [sveltekit()]
};

export default config;