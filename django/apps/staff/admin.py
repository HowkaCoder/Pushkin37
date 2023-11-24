from django.contrib.admin import *

from .models import Staff

@register(Staff)
class StaffAdmin(ModelAdmin):

    list_display = ('id', 'fullname', 'position')
    list_display_links = ('id', 'fullname', 'position')