#!/usr/bin/env python
# coding: utf-8

# start project: django-admin.py startproject placeholder --template=project_name

import os
import sys

from django.conf import settings

DEBUG = os.environ.get('DEBUG', 'on') == 'on'
SECRET_KEY = os.environ.get('SECRET_KEY', 'pv0^7c$a^u(wp+-3ksswb0k0=@tuv_64&7y@!rni#ub-q642d=')
ALLOWED_HOSTS = os.environ.get('ALLOWED_HOSTS', 'localhost').split(',')


# configure the settings before makeing any additional imports from Django.
settings.configure(
    DEBUG=DEBUG,
    SECRET_KEY=SECRET_KEY,
    ROOT_URLCONF=__name__,
    ALLOWED_HOSTS=ALLOWED_HOSTS,
    MIDDLEWARE_CLASSES=(
        'django.middleware.common.CommonMiddleware',
        'django.middleware.csrf.CsrfViewMiddleware',
        'django.middleware.clickjacking.XFrameOptionsMiddleware'
    )
)


from django.conf.urls import url
from django.core.wsgi import get_wsgi_application
from django.http import HttpResponse, HttpResponseBadRequest
from django import forms


class ImageForm(forms.Form):
    width = forms.IntegerField(min_value=1, max_value=2000)
    height = forms.IntegerField(min_value=1, max_value=2000)


def placeholder(request, width, height):
    # TODO: Rest of the view will go here
    form = ImageForm({'width': width, 'height': height})
    if form.is_valid():
        return HttpResponse("OK")
    else:
        return HttpResponseBadRequest('Invalid Image Request.')


def index():
    return HttpResponse('Home')

urlpatterns = (
    url(r'^image/(?P<width>[0-9]+)x(?P<height>[0-9]+)/$', placeholder, name='placeholder'),
    url(r'^$', index, name='homepage'),
)


application = get_wsgi_application()

if __name__ == "__main__":
    from django.core.management import execute_from_command_line
    execute_from_command_line(sys.argv)
