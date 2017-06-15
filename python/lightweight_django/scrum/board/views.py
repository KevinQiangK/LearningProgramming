from rest_framework import authentication, permissions, viewsets
from .models import Sprint
from .serializers import SprintSerial

# Create your views here.

class DefaultMixin(object):
    authentication_classes = (
        authentication.BaseAuthentication,
        authentication.TokenAuthentication
    )

    permission_classes = (
        permissions.is_authenticated,
    )

    paginate_by = 25
    paginate_by_param = 'page_size'
    max_paginate_by = 100


class SprintViewSet(viewsets.ModelViewSet):
    queryset = Sprint.objects.order_by('end')
    serializer_class = SprintSerial

