from liferaybeat import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Liferaybeat normally
        """
        self.render_config_template(
                path=os.path.abspath(self.working_dir) + "/log/*"
        )

        liferaybeat_proc = self.start_beat()
        self.wait_until( lambda: self.log_contains("liferaybeat is running"))
        exit_code = liferaybeat_proc.kill_and_wait()
        assert exit_code == 0
