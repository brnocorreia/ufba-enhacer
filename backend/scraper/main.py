import time
from selenium.webdriver.support import expected_conditions as EC

from src.interfaces.driver import Driver
from src.factory.disciplines.disciplines import Disciplines
from src.utils.xpaths import Login as login

def main():
    driver = Driver()
    
    driver.get(login.URL)
    cpf = driver.find_element(login.CPF, 10, EC.element_to_be_clickable)
    cpf.send_keys('CPF AQUI')

    password = driver.find_element(login.PASS, 10, EC.element_to_be_clickable)
    password.send_keys('SENHA AQUI')

    driver.find_element(login.LOGIN_BUTTON, 10, EC.element_to_be_clickable).click()

    disciplines = Disciplines(driver)
    disciplines.get_all_disciplines()

if __name__ == '__main__':
    main()
