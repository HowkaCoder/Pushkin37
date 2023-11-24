from django.contrib.admin import *

from .models import StaffPosition

@register(StaffPosition)
class StaffPositionAdmin(ModelAdmin):

    list_display = ('id', 'position')
    list_display_links = ('id', 'position')