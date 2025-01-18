import smtplib
from email.mime.text import MIMEText
from email.utils import formataddr
import ssl
import logging
import contextlib
from typing import List
import time


class EmailSender:
    def __init__(self, username: str, password: str, hostname: str, nickname: str = ""):
        self.username = username  # 邮箱用户名
        self.password = password  # 邮箱密码
        self.hostname = hostname  # SMTP 主机地址
        self.nickname = nickname or username  # 邮件发送者昵称

    def send_email(self, subject: str, body: str, recipients: List[str], timeout: int = 60) -> None:
        """
        使用 SMTP 通过 STARTTLS 发送电子邮件。

        :param subject: 邮件主题
        :param body: 邮件正文（支持 HTML 格式）
        :param recipients: 收件人邮箱地址列表
        :param timeout: 操作超时时间（秒）
        """
        # 创建邮件内容，指定为 HTML 格式
        msg = MIMEText(body, "html", "utf-8")
        msg["From"] = formataddr((self.nickname, self.username))
        msg["Subject"] = subject
        msg["To"] = ", ".join(recipients)  # 设置所有收件人

        context = ssl.create_default_context()

        try:
            # 连接到 SMTP 服务器
            with contextlib.closing(smtplib.SMTP(self.hostname, 587)) as server:
                server.starttls(context=context)  # 启用 STARTTLS
                server.login(self.username, self.password)  # 登录 SMTP 服务器

                # 发送邮件并计时
                start_time = time.time()
                logging.info(f"Sending emails form {self.username} to {recipients}")
                server.sendmail(self.username, recipients, msg.as_string())

                elapsed_time = time.time() - start_time
                if elapsed_time > timeout:
                    logging.error("Emails sending timeout.")
                    raise TimeoutError("Emails sending timeout.")

                logging.info(f"Emails successfully sent to {recipients}")

        except Exception as e:
            logging.error(f"Emails sending failed: {e}")
            raise

    def send_email_tls(self, subject: str, body: str, recipients: List[str], timeout: int = 60) -> None:
        """
        使用 SMTP 通过 SSL 发送电子邮件。

        :param subject: 邮件主题
        :param body: 邮件正文（支持 HTML 格式）
        :param recipients: 收件人邮箱地址列表
        :param timeout: 操作超时时间（秒）
        """
        # 创建邮件内容，指定为 HTML 格式
        msg = MIMEText(body, "html", "utf-8")
        msg["From"] = formataddr((self.nickname, self.username))
        msg["Subject"] = subject
        msg["To"] = ", ".join(recipients)  # 设置所有收件人

        context = ssl.create_default_context()

        try:
            # 连接到 SMTP 服务器
            with smtplib.SMTP_SSL(self.hostname, 465, context=context) as server:
                server.login(self.username, self.password)  # 登录 SMTP 服务器

                # 发送邮件并计时
                start_time = time.time()
                logging.info(f"Sending emails From {self.username} to {recipients}")
                server.sendmail(self.username, recipients, msg.as_string())

                elapsed_time = time.time() - start_time
                if elapsed_time > timeout:
                    logging.error("Emails sending timeout")
                    raise TimeoutError("Emails sending timeout")

                logging.info(f"Emails successfully sent to {recipients}")

        except Exception as e:
            logging.error(f"Emails sending failed: {e}")
            raise
