from rest_framework.status import *
from rest_framework.response import Response
from rest_framework.viewsets import ViewSet

from apps.pride_students.models import PrideStudent
from apps.staff_position.models import StaffPosition
from apps.staff.models import Staff

from .serializers import (
    PrideStudentSerializer,
    StaffPositionSerializer,
    StaffSerializer
)

class StaffPositionViewSet(ViewSet):

    def list(self, request):
        queryset = StaffPosition.objects.all()
        serializer = StaffPositionSerializer(queryset, many=True)
        return Response(
            serializer.data,
            status=HTTP_200_OK
        )

    def create(self, request):
        serializer = StaffPositionSerializer(request.data)
        if serializer.is_valid():
            serializer.save()
            return Response(
                serializer.data,
                status=HTTP_201_CREATED
            )
        return Response(
            serializer.errors,
            status=HTTP_400_BAD_REQUEST
        )
    
    def retrieve(self, request, pk=None):
        queryset = StaffPosition.objects.get(pk=pk)
        serializer = StaffPositionSerializer(queryset)
        return Response(
            serializer.data,
            status=HTTP_200_OK
        )
    
    def partial_update(self, request, pk=None):
        queryset = StaffPosition.objects.get(pk=pk)
        serializer = StaffPositionSerializer(queryset, data=request.data, partial=True)
        if serializer.is_valid():
            serializer.save()
            return Response(
                serializer.data,
                status=HTTP_202_ACCEPTED
            )
        return Response(
            serializer.errors,
            status=HTTP_400_BAD_REQUEST
        )
    
    def destroy(self, request, pk=None):
        queryset = StaffPosition.objects.get(pk=pk)
        queryset.delete()
        return Response(
            'Deleted',
            status=HTTP_204_NO_CONTENT
        )
    
class StaffViewSet(ViewSet):

    def list(self, request):
        queryset = Staff.objects.all()
        serializer = StaffSerializer(queryset, many=True, context={'request': request})
        return Response(
            serializer.data,
            status=HTTP_200_OK
        )

    def create(self, request):
        serializer = StaffSerializer(data=request.data, context={'request': request})
        if serializer.is_valid():
            serializer.save()
            return Response(
                serializer.data,
                status=HTTP_201_CREATED
            )
        return Response(
            serializer.errors,
            status=HTTP_400_BAD_REQUEST
        )
    
    def retrieve(self, request, pk=None):
        queryset = Staff.objects.get(pk=pk)
        serializer = StaffSerializer(queryset, context={'request': request})
        return Response(
            serializer.data,
            status=HTTP_200_OK
        )
    
    def partial_update(self, request, pk=None):
        queryset = Staff.objects.get(pk=pk)
        serializer = StaffSerializer(
            queryset, 
            data=request.data, 
            partial=True, 
            context={'request': request}
        )
        if serializer.is_valid():
            serializer.save()
            return Response(
                serializer.data,
                status=HTTP_202_ACCEPTED
            )
        return Response(
            serializer.errors,
            status=HTTP_400_BAD_REQUEST
        )
    
    def destroy(self, request, pk=None):
        queryset = Staff.objects.get(pk=pk)
        queryset.delete()
        return Response(
            'Deleted',
            status=HTTP_204_NO_CONTENT
        )
    
class PrideStudentViewSet(ViewSet):

    def list(self, request):
        queryset = PrideStudent.objects.all()
        serializer = PrideStudentSerializer(queryset, many=True, context={'request': request})
        return Response(
            serializer.data,
            status=HTTP_200_OK
        )

    def create(self, request):
        serializer = PrideStudentSerializer(data=request.data, context={'request': request})
        if serializer.is_valid():
            serializer.save()
            return Response(
                serializer.data,
                status=HTTP_201_CREATED
            )
        return Response(
            serializer.errors,
            status=HTTP_400_BAD_REQUEST
        )
    
    def retrieve(self, request, pk=None):
        queryset = PrideStudent.objects.get(pk=pk)
        serializer = PrideStudentSerializer(queryset, context={'request': request})
        return Response(
            serializer.data,
            status=HTTP_200_OK
        )
    
    def partial_update(self, request, pk=None):
        queryset = PrideStudent.objects.get(pk=pk)
        serializer = PrideStudentSerializer(
            queryset, 
            data=request.data, 
            partial=True, 
            context={'request': request}
        )
        if serializer.is_valid():
            serializer.save()
            return Response(
                serializer.data,
                status=HTTP_202_ACCEPTED
            )
        return Response(
            serializer.errors,
            status=HTTP_400_BAD_REQUEST
        )
    
    def destroy(self, request, pk=None):
        queryset = PrideStudent.objects.get(pk=pk)
        queryset.delete()
        return Response(
            'Deleted',
            status=HTTP_204_NO_CONTENT
        )