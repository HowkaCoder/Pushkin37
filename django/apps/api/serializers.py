from rest_framework.serializers import *

from apps.staff.models import Staff
from apps.pride_students.models import PrideStudent
from apps.staff_position.models import StaffPosition

class StaffPositionSerializer(ModelSerializer):

    class Meta:

        model = StaffPosition
        fields = '__all__'

class StaffSerializer(ModelSerializer):

    class Meta:

        model = Staff
        fields = '__all__'

class PrideStudentSerializer(ModelSerializer):

    class Meta:

        model = PrideStudent
        fields = '__all__'