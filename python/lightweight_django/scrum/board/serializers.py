from rest_framework import serializers
from .models import Sprint

class SprintSerial(serializers.ModelSerializer):
    class Meta:
        model = Sprint
        field = ('id', 'name', 'description', 'end',)
