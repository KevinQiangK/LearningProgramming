import sys
from django.conf import settings

settings.configure(
    DEBUG=True,
    SECRET_KEY='pv0^7c$a^u(wp+-3ksswb0k0=@tuv_64&7y@!rni#ub-q642d=',
    ROOT_URLCONF='sitebuilder.urls',
    MIDDLEWARE_CLASS=(),
    INSTALLED_APPS=(
        'django.contrib.staticfiles',
        'sitebuilder'
    ),
    STATIC_URL='/static/'
)

if __name__ == "__main__":
    from django.core.management import execute_from_command_line
    execute_from_command_line(sys.argv)