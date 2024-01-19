import { test } from '@playwright/test';
import * as fs from 'fs';

const config: {  username: string; password: string } = JSON.parse(fs.readFileSync('.auth/config.json', 'utf-8'));


test('login mitch and see email address in header', async ({ page }) => {
  await page.goto('/login');

  // Fill in the login form
  const emailElement = page.locator('input[name="email"]');
  await emailElement.fill(config.login.username);

  const passElement = page.locator('input[name="password"]');
  await passElement.fill(config.login.password);


  // Submit the login form
  await page.click('button[type="submit"]'); // Adjust selector for the login button

  // Wait for a page with a specific header to load
  let response = await page.waitForResponse(response => response.status() === 200);

  console.log('Email header value:', response);
  await page.screenshot({ path: './custom-folder/screenshot.png' });
});

