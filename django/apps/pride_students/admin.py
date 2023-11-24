from django.contrib.admin import *

from .models import PrideStudent

@register(PrideStudent)
class PrideStudentAdmin(ModelAdmin):

    list_display = ('id', 'fullname', 'graduation_year')
    list_display_links = ('id', 'fullname', 'graduation_year')