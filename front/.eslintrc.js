module.exports = {
    root: true,
    extends: ['plugin:@typescript-eslint/recommended','eslint:recommended', 'plugin:svelte/recommended', 'prettier'],
    parserOptions: {
        sourceType: 'module',
        ecmaVersion: 2020,
        extraFileExtensions: ['.svelte']
    },
    env: {
        browser: true,
        es2017: true,
        node: true
    }
};

