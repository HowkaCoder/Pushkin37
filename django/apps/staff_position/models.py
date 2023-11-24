from django.db.models import *

class StaffPosition(Model):

    position = CharField(
        'Позиция',
        max_length=128
    )

    def __str__(self):
        return f'{self.position}'
    
    class Meta:

        verbose_name = 'Позиция сотрудника'
        verbose_name_plural = 'Позиции сотрудников'