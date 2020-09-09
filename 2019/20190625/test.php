case ZBX_DB_MYSQL:
                                $DB['DB'] = @mysqli_connect($DB['SERVER'], $DB['USER'], $DB['PASSWORD'], $DB['DATABASE'], $DB['PORT']);
                                mysqli_set_charset($DB['DB'],'utf8mb4');
                                $DB['DB']->query("SET NAMES utf8mb4 COLLATE utf8mb4_unicode_ci");
                                if (!$DB['DB']) {
                                        $error = 'Error connecting to database: '.trim(mysqli_connect_error());
                                        $result = false;
                                }
                                elseif (mysqli_autocommit($DB['DB'], true) === false) {
                                        $error = 'Error setting auto commit.';
                                        $result = false;
                                }
                                else {
                                        DBexecute('SET NAMES utf8mb4');
                                }

                                if ($result) {
                                        $dbBackend = new MysqlDbBackend();
                                }
                                DBexecute('SET NAMES utf8mb4');
                                break;