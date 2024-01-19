import { test as base } from '@playwright/test';

export const test = base.extend({
  // Override storageState fixture to authenticate before each test
  async storageState(context) { // Add the async keyword here
    const browser = await context.newPage();
    await browser.authenticate({
      // Specify the path to your .auth file
      auth: 'playwright/auth.json'
    });
    await browser.close();
    return browser.storageState();
  },
});