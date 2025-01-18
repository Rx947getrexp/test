from datetime import datetime, timedelta
import logging
import time
import email_service
import log

from db_util import Speed  # 导入 Speed 类

# 邮件发送配置
batch_size=99 #每次发送的收件人数量上限100人每封
daily_limit=499 #每天发送邮件总数上限500封

def get_email_content():
    # 邮件标题和正文
    subject = "Подарок для вас – 15 дней бесплатного использования!"
    body = "<html>"
    body += "<body>"
    body += "<p>Дорогой пользователь!</p>"
    body += "<p>Мы заметили, что вы давно не пользовались нашим продуктом, и хотели бы напомнить вам о преимуществах, которые вы упускаете. В знак благодарности за ваше доверие и выбор нашего сервиса мы предоставляем вам 15 дней бесплатного использования.</p>"
    body += "<p>Если вы уже удалили наше приложение, вы всегда можете скачать его снова по следующей ссылке: <a href=\"https://www.yyy360.xyz\">[Герой VPN]https://www.yyy360.xyz</a></p>"
    body += "<p>Ваши 15 бесплатных дней уже активированы и ждут вас!</p>"
    body += "<p>Если у вас возникнут какие-либо вопросы или проблемы, наша команда поддержки всегда готова помочь. Мы будем рады снова видеть вас среди наших пользователей!</p>"
    body += "<p>С уважением,  </p>"
    body += "<p>Команда [Герой VPN]</p>"
    body += "</body>"
    body += "</html>"
    # 返回邮件标题和正文
    return subject, body

def send_bulk_emails(sender, subject, body, recipients):
    """
    分批次发送邮件，遵循服务商限制。
    
    :param sender: EmailSender 对象
    :param subject: 邮件主题
    :param body: 邮件正文
    :param recipients: 收件人列表
    :param batch_size: 每批次的收件人数量限制（默认100）
    :param daily_limit: 每日邮件总数限制（默认500）
    """
    total_recipients = len(recipients)
    total_batches = (total_recipients + batch_size - 1) // batch_size
    sent_count = 0  # 当前已发送的邮件计数

    for batch_index in range(total_batches):
        # 获取当前批次的收件人
        start = batch_index * batch_size
        end = min(start + batch_size, total_recipients)
        current_batch = recipients[start:end]

        try:
            # 检查是否达到每日限制
            if sent_count >= daily_limit:
                # logging.info(f"每日发送限制已达 {daily_limit} 封，程序暂停至次日...")
                logging.info(f"Daily sending limit of {daily_limit} emails reached, the program will pause until the next day...")
                time.sleep(24 * 60 * 60)  # 暂停 24 小时

            # 发送当前批次的邮件
            # logging.info(f"正在发送第 {batch_index + 1}/{total_batches} 批次...")
            logging.info(f"Sending batch {batch_index + 1}/{total_batches}...")

            sender.send_email_tls(subject, body, current_batch)

            sent_count += 1  # 增加已发送邮件计数
            # logging.info(f"成功发送第 {batch_index + 1} 批次邮件，当前已发送 {sent_count} 封邮件。")
            logging.info(f"Successfully sent batch {batch_index + 1} emails, a total of {sent_count} emails sent so far.")


        except Exception as e:
            # logging.error(f"发送第 {batch_index + 1} 批次邮件时发生错误: {e}")
            logging.error(f"An error occurred while sending batch {batch_index + 1} emails: {e}")
            continue  # 跳过错误，继续发送后续批次

        # 如果需要，可以设置发送间隔，避免触发反垃圾邮件机制
        time.sleep(60)

# 获取需要发送的用户邮件地址
def get_recovery_users():
    # 获取当前时间
    now = datetime.now()
    # 当前时间减去 30 天
    thirty_days_ago = now - timedelta(days=30)
    # 获取 Unix 时间戳
    thirty_days_ago_timestamp = int(thirty_days_ago.timestamp())
    # 获取需要发送的用户邮件
    recipients = Speed().get_recovery_users(thirty_days_ago, thirty_days_ago_timestamp)  # 获取需要发送的用户邮件
    return recipients

def run():
    # 开始执行任务
    logging.info("=" * 10 + "sending emails" + "=" * 10)
    # 邮件发送配置信息
    sender = email_service.EmailSender(username="heronet@heronet.shop", password="pingguoqm23", hostname="smtpout.secureserver.net", nickname="Герой VPN")
    # 获取邮件标题和正文
    subject, body = get_email_content()
    # 获取需要发送的用户邮件
    recipients = get_recovery_users()
    # recipients = ["pmm73219@gmail.com", "273768414@qq.com", "shenfuqing@163.com", "mmp73219@outlook.com"] #测试发送
    # 分批次发送邮件，遵循服务商限制，使用 TLS
    send_bulk_emails(sender, subject, body, recipients)

# 示例用法
if __name__ == "__main__":
    log.init_logging("./log/user_recovery.log")
    try:
        run()
    except Exception as e:
        # logging.error(f"发生错误: {e}")
        logging.error(f"An error occurred: {e}")