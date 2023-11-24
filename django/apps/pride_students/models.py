from django.db.models import *

class PrideStudent(Model):

    fullname = CharField(
        verbose_name='ФИО',
        max_length=128
    )

    graduation_year = PositiveIntegerField(
        verbose_name='Год выпуска'
    )

    description = TextField(
        verbose_name='Информация об ученике'
    )

    photo = ImageField(
        verbose_name='Фото',
        upload_to='pride_students_photos'
    )

    current_position = CharField(
        verbose_name='Текущяя позиция',
        max_length=128
    )

    def __str__(self):
        return f'{self.fullname}'
    
    class Meta:

        verbose_name = 'Гордость школы'
        verbose_name_plural = 'Гордости школы'