from django.db.models import *

from apps.staff_position.models import StaffPosition

class Staff(Model):

    fullname = CharField(
        verbose_name='ФИО',
        max_length=128
    )

    position = ForeignKey(
        StaffPosition,
        verbose_name='Позиция сотрудника',
        on_delete=CASCADE
    )

    description = TextField(
        verbose_name='О сотруднике',
    )

    photo = ImageField(
        verbose_name='Фото сотрудника',
        upload_to='staff_photos'
    )

    experience = PositiveIntegerField(
        verbose_name='Стаж сотрудника (в годах)'
    )

    def __str__(self):
        return f'{self.fullname}'
    
    class Meta:

        verbose_name = 'Сотрудник'
        verbose_name_plural = 'Сотрудники'