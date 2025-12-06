import time

from locust import HttpUser, task, between

class UserAPITest(HttpUser):
    wait_time = between(1, 2)
    token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdWIiOjEsIlJvbGUiOiJhZG1pbiIsImV4cCI6MTc2NTExNTAxMSwiaWF0IjoxNzY1MDI4NjExfQ.QYl59RcMyqE6mLnMPx3DFKP0Tl6hvYQgnWQPXEfit8g"
    def on_start(self):
        self.headers = {
            "Authorization": f"{self.token}",
            "Content-Type": "application/json"
        }

    # @task(3)
    # def get_users(self):
    #     self.client.get(
    #         "/api/v1/users/",
    #         headers=self.headers,
    #         name="GET /api/v1/users/"
    #     )

    # @task(1)
    # def get_user_by_id(self):
    #     user_id = 1
    #     self.client.get(
    #         f"/api/v1/users/{user_id}",
    #         headers=self.headers,
    #         name="GET /api/v1/users/:id"
    #     )
    #
    @task(2)
    def create_user(self):
        timestamp = int(time.time() * 1000000)
        username = f"test{timestamp}@example.com"

        payload = {
            "email": username,
            "password": "securepass123"
        }

        self.client.post(
            "/api/v1/auth/register",
            json=payload,
            headers=self.headers,
            name="POST /api/v1/auth/register"
        )