from datetime import datetime
import logging
import time
import email_service
import log
from db_util import SpeedReport

# 邮件发送配置
BATCH_SIZE = 98  # 每批次真实用户数（确保 BATCH_SIZE + 测试邮箱数量 ≤ 100）
DAILY_LIMIT = 500  # 每日发送上限（包含测试邮箱）
MAX_RETRIES = 3  # 失败重试次数
RETRY_DELAY = 60  # 失败重试等待时间（秒）
BATCH_INTERVAL = 60 * 30  # 每批次间隔 30 分钟，防止触发反垃圾机制
TEST_EMAILS = [  # 测试邮箱列表（需计入限额）
    "273768414@qq.com",
    "shenfuqing@163.com"
]

def get_email_content():
    """
    获取邮件标题和正文
    """
    subject = "Подарок для вас – 15 дней бесплатного использования!"
    body = """
    <html>
    <body>
        <p>Дорогой пользователь!</p>
        <p>Мы заметили, что вы давно не пользовались нашим продуктом, и хотели бы напомнить вам о преимуществах, которые вы упускаете. В знак благодарности за ваше доверие и выбор нашего сервиса мы предоставляем вам 15 дней бесплатного использования.</p>
        <p>Если вы уже удалили наше приложение, вы всегда можете скачать его снова по следующей ссылке: <a href="https://www.yyy360.xyz">[Герой VPN]https://www.yyy360.xyz</a></p>
        <p>Ваши 15 бесплатных дней уже активированы и ждут вас!</p>
        <p>Если у вас возникнут какие-либо вопросы или проблемы, наша команда поддержки всегда готова помочь. Мы будем рады снова видеть вас среди наших пользователей!</p>
        <p>С уважением,</p>
        <p>Команда [Герой VPN]</p>
    </body>
    </html>
    """
    return subject, body

def get_recovery_users():
    """获取需要发送挽回邮件的用户列表，每天最多获取 500 个"""
    return SpeedReport().get_unsent_recovery_users(limit=490)  # ✅ 获取 500 个（含 10 测试）

def update_send_emails_status(emails):
    """更新邮件发送状态"""
    if not emails:
        return
    try:
        email_list_str = "'" + "','".join(emails) + "'"
        SpeedReport().update_recovery_emails_status(email_list_str)
    except Exception as e:
        logging.error(f"更新数据库状态失败：{str(e)}")

def send_bulk_emails(sender, subject, body, recipients):
    """
    分批次发送邮件，每批 100 封（98 个真实用户 + 2 个测试邮箱），每天最多 5 批次（共 500 封）
    """
    total_recipients = min(len(recipients), 500)  # ✅ 限制真实用户最多 500
    total_batches = min((total_recipients + 97) // 98, 5)  # ✅ 最多 5 批次
    sent_count = 0  # 记录已发送数量（包含测试邮箱）

    for batch_index in range(total_batches):
        start = batch_index * 98
        end = min(start + 98, total_recipients)
        current_main_emails = recipients[start:end]

        current_batch = current_main_emails + TEST_EMAILS  # ✅ 每批加 2 个测试邮箱
        batch_size = len(current_batch)

        logging.info(f"正在发送第 {batch_index+1}/{total_batches} 批次：{batch_size} 封邮件")

        # 发送重试逻辑
        retry_count = 0
        while retry_count < MAX_RETRIES:
            try:
                sender.send_email_tls(subject, body, current_batch)
                sent_count += batch_size  

                logging.info(f"成功发送 {batch_size} 封邮件，累计已发送：{sent_count}/500")

                update_send_emails_status(current_main_emails)  # ✅ 更新数据库
                break
            except Exception as e:
                logging.error(f"第 {batch_index+1} 批次发送失败（第 {retry_count+1} 次重试），错误信息：{str(e)}")
                retry_count += 1
                if retry_count < MAX_RETRIES:
                    time.sleep(RETRY_DELAY)
                else:
                    logging.error(f"第 {batch_index+1} 批次在 {MAX_RETRIES} 次尝试后仍发送失败")

        if sent_count >= 500:
            logging.info("已达到每日上限 500，停止后续发送")
            break

        logging.info(f"等待 {BATCH_INTERVAL} 秒后发送下一批次...")
        time.sleep(BATCH_INTERVAL)  # ✅ 控制间隔

def run():
    """主运行函数"""
    logging.info("========== 开始执行用户挽回邮件发送任务 ==========")

    # 初始化邮件服务
    sender = email_service.EmailSender(
        username="heronet@heronet.shop",
        password="pingguoqm23",
        hostname="smtpout.secureserver.net",
        nickname="Герой VPN"
    )

    # 获取邮件内容和收件人
    subject, body = get_email_content()
    recipients = get_recovery_users()

    if not recipients:
        logging.info("未找到需要发送邮件的用户")
        return

    send_bulk_emails(sender, subject, body, recipients)
    logging.info("邮件发送任务执行完毕")

if __name__ == "__main__":
    log.init_logging("/shell/report/log/user_recovery.log")
    # log.init_logging("./log/user_recovery.log")

    try:
        run()
    except Exception as e:
        logging.critical(f"程序发生严重错误：{str(e)}", exc_info=True)
