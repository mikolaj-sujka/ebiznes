from selenium.webdriver.support import expected_conditions as EC
from utils.base_test import BaseTest
from selenium.webdriver.common.by import By

class TestQodeca(BaseTest):
    def test_homepage_title(self):
        self.driver.get('https://www.qodeca.com')
        self.assertIn('Qodeca', self.driver.title)

    def test_homepage_elements(self):
        self.driver.get('https://www.qodeca.com')
        header = self.wait.until(EC.presence_of_element_located((By.TAG_NAME, 'header')))
        self.assertTrue(header.is_displayed())

        footer = self.wait.until(EC.presence_of_element_located((By.TAG_NAME, 'footer')))
        self.assertTrue(footer.is_displayed())

    def test_navigation_links(self):
        self.driver.get('https://www.qodeca.com')
        nav_links = self.driver.find_elements(By.CSS_SELECTOR, 'nav a')
        self.assertGreater(len(nav_links), 0)
        for link in nav_links:
            self.assertTrue(link.is_displayed())

    def test_nav_clients_link(self):
        self.driver.get('https://www.qodeca.com')
        clients_link = self.wait.until(EC.element_to_be_clickable((By.LINK_TEXT, 'Clients')))
        self.assertTrue(clients_link.is_displayed())
        clients_link.click()

    def test_nav_services_link(self):
        self.driver.get('https://www.qodeca.com')
        services_link = self.wait.until(EC.element_to_be_clickable((By.LINK_TEXT, 'Services')))
        self.assertTrue(services_link.is_displayed())
        services_link.click()

    def test_nav_about_us_link(self):
        self.driver.get('https://www.qodeca.com')
        about_us_link = self.wait.until(EC.element_to_be_clickable((By.LINK_TEXT, 'About Us')))
        self.assertTrue(about_us_link.is_displayed())
        about_us_link.click()

    def test_nav_insights_link(self):
        self.driver.get('https://www.qodeca.com')
        insights_link = self.wait.until(EC.element_to_be_clickable((By.LINK_TEXT, 'Insights')))
        self.assertTrue(insights_link.is_displayed())
        insights_link.click()