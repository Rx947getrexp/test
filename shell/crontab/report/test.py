
import user
import util

def run():
    t = util.Time(util.get_yesterday_date())
    m =util.Time(util.get_month_date())
    user.ReportUser(t.date, t.get_start_time(), t.get_end_time(), m.get_start_time()).report_ad_daily()

if __name__ == '__main__':
    try:
        run()
    except Exception as e:
        print(f'An error occurred: {e}')