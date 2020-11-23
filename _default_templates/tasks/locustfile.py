from locust import HttpUser, TaskSet, task, between, events, constant
from locust.exception import StopUser

DEBUG_MODE = False


def print_log(url, response):
    if not DEBUG_MODE:
        return

    if response.status_code in [200, 201]:
        print(f'[{response.status_code}] {url}')
    else:
        print(f'[{response.status_code}] {url} {response.text}')


class WebsiteTasks(TaskSet):

    def __init__(self, parent):
        super().__init__(parent)
        self.default_headers = None

    # call with starting new test
    @events.test_start.add_listener
    def on_test_start(**kwargs):
        return

    # call with stopping new test
    @events.test_stop.add_listener
    def on_test_stop(**kwargs):
        return

    # get, post, put, delete helper methods
    def get(self, url, headers=None, name=None):
        response = self.client.get(url, headers=headers or self.default_headers, name=name or url)
        print_log(url, response)
        return response

    def post(self, url, json_body, headers=None, name=None):
        response = self.client.post(url, json=json_body, headers=headers or self.default_headers, name=name or url)
        print_log(url, response)
        return response

    def put(self, url, json_body, headers=None, name=None):
        response = self.client.put(url, json=json_body, headers=headers or self.default_headers, name=name or url)
        print_log(url, response)
        return response

    def delete(self, url, json_body, headers=None, name=None):
        response = self.client.delete(url, json=json_body, headers=headers or self.default_headers, name=name or url)
        print_log(url, response)
        return response

    # login example
    def login(self, email, password):
        url = "/api/v1/login"
        headers = {'Content-Type': "application/json"}
        body = {
            "email": email,
            "password": password,
        }
        response = self.post(url, body, headers)

        if response.status_code != 200:
            print(f'StopUser: Login Request Failed: {url}, {email}, {response.status_code}, {response.text}')
            raise StopUser()

        token = response.json()["token"]
        self.default_headers = {
            'Content-Type': "application/json",
            "Authorization": "Bearer " + token
        }

    # logout example
    def logout(self):
        self.post("/api/v1/logout", {})

    # call with starting new task
    def on_start(self):
        # self.login("email", "password")
        return

    # call with stopping new task
    def on_stop(self):
        # self.logout()
        return

    ######################################################################
    # write your tasks ###################################################
    ######################################################################
    @task(60)
    def index(self):
        self.get("/")

    @task(40)
    def health(self):
        self.get("/health")


class WebsiteUser(HttpUser):
    tasks = [WebsiteTasks]

    # If you want no wait time between tasks
    # wait_time = constant(0)
    wait_time = between(1, 2)

    # default target host
    host = "http://www.example.com"
