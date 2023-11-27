from django.urls import path
from rest_framework.routers import DefaultRouter

from .views import (
    StaffPositionViewSet,
    StaffViewSet,
    PrideStudentViewSet
)

router = DefaultRouter()
router.register('staff-position', StaffPositionViewSet, basename='staff-position')
router.register('staff', StaffViewSet, basename='staff')
router.register('pride-student', PrideStudentViewSet, basename='pride-student')

urlpatterns = []
urlpatterns += router.urls