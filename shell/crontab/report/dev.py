
import user
import util

import logging
import log

TASK_NAME = 'dev_py'

def run():
    t = util.Time(util.get_yesterday_date())
    m = util.Time(util.get_month_date())
    user.ReportUser(t.date, t.get_start_time(), t.get_end_time(), m.get_start_time()).report_channel_retaind_daily()

if __name__ == '__main__':
    log.init_logging("./log/%s.log" % TASK_NAME)
    logging.info("\n\n\n start %s" % TASK_NAME)
    try:
        run()
    except Exception as e:
        print(f'An error occurred: {e}')