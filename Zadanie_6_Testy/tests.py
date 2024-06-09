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

class TestGitHub(BaseTest):
    def test_homepage_title(self):
        self.driver.get('https://github.com')
        self.assertIn('GitHub', self.driver.title)

    def test_homepage_elements(self):
        self.driver.get('https://github.com')
        header = self.wait.until(EC.presence_of_element_located((By.TAG_NAME, 'header')))
        self.assertTrue(header.is_displayed())

        footer = self.wait.until(EC.presence_of_element_located((By.TAG_NAME, 'footer')))
        self.assertTrue(footer.is_displayed())

    def test_navigation_links(self):
        self.driver.get('https://github.com')
        nav_links = self.driver.find_elements(By.CSS_SELECTOR, 'nav a')
        self.assertGreater(len(nav_links), 0)
        for link in nav_links:
            self.assertFalse(link.is_displayed())

    def test_footer_links(self):
        self.driver.get('https://github.com')
        footer_links = self.driver.find_elements(By.CSS_SELECTOR, 'footer a')
        self.assertGreater(len(footer_links), 0)
        for link in footer_links:
            self.assertTrue(link.is_displayed())

class TestXKom(BaseTest):
    def test_homepage_title(self):
        self.driver.get('https://www.x-kom.pl')
        self.assertIn('x-kom', self.driver.title)

    def test_navigation_bar(self):
        self.driver.get('https://www.x-kom.pl')
        nav_bar = self.wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, 'nav')))
        self.assertTrue(nav_bar.is_displayed())

    def test_search_functionality(self):
        self.driver.get('https://www.x-kom.pl')
        search_box = self.wait.until(EC.presence_of_element_located((By.NAME, 'q')))
        search_box.send_keys('laptop')
        search_box.submit()

        results = self.wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, '.sc-6n68ef-0')))
        self.assertTrue(results.is_displayed())

    def test_footer_elements(self):
        self.driver.get('https://www.x-kom.pl')
        footer = self.wait.until(EC.presence_of_element_located((By.TAG_NAME, 'footer')))
        self.assertTrue(footer.is_displayed())

        footer_links = footer.find_elements(By.CSS_SELECTOR, 'a')
        self.assertGreater(len(footer_links), 0)
        for link in footer_links:
            self.assertTrue(link.is_displayed())

    def test_category_menu(self):
        self.driver.get('https://www.x-kom.pl')
        category_menu = self.wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, '.sc-1v4pze9-1')))
        self.assertTrue(category_menu.is_displayed())

    def test_cart_button(self):
        self.driver.get('https://www.x-kom.pl')
        cart_button = self.wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR, 'a[href="/koszyk"]')))
        self.assertTrue(cart_button.is_displayed())
        cart_button.click()

        self.assertIn('Koszyk', self.driver.title)

    def test_promotions_link(self):
        self.driver.get('https://www.x-kom.pl')
        promotions_link = self.wait.until(EC.element_to_be_clickable((By.LINK_TEXT, 'Promocje')))
        self.assertTrue(promotions_link.is_displayed())
        promotions_link.click()

        self.assertIn('Promocje', self.driver.title)

    def test_contact_page(self):
        self.driver.get('https://www.x-kom.pl')
        contact_link = self.wait.until(EC.element_to_be_clickable((By.LINK_TEXT, 'Kontakt')))
        self.assertTrue(contact_link.is_displayed())
        contact_link.click()

        self.assertIn('Kontakt', self.driver.title)