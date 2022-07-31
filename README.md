# parallelism-analyzer

### Назначение проекта

Проект предназначен для реализации инструмента по анализу времени выполнения параллельных реализаций алгоритмов

### Структура проекта

- _**assets**_ - ресурсные файлы, используемые в тестировании алгоритмов
- _**cmd**_ - директория с *main* файлом
- _**internal**_ - внутренние пакеты проекта
    - *internal/algorithms* - алгоритмы на примере которых производилось тестирование
        - *internal/algorithms/dithering* - алгоритмы псевдотонирования изображений
        - *internal/algorithms/fourier*  - алгоритмы Дискретного преобразования Фурье
- _**pkg**_ - основные пакеты проекта
    - *pkg/analyzer* - реализация анализатора
- _**results**_ - временная директория, хранящая результаты unit-тестов, генерирующих файлы
- _**test**_ - временная директория, хранящая результаты unit-тестирования с флагом -cover

### Описание пакетов
- _**internal/algorithms/dithering**_
    - _**internal/algorithms/dithering/dithering.go**_
        - *func ditheringMatrix(order int) (result [][]int, errF error)* - функция рассчета матриц размытия (параметр *order* должен быть степенью 2 больше 1)
        - *func readSourceImg(srcImgPath string) (img image.Image, errF error)* - функция чтения в память изображения, находящегося по пути, указанному в параметре *srcImgPath*
        - *func ThresholdDithering(srcImgPath string, dstImgPath string, threshold int) (errF error)* - функция, реализующая алгоритм усечения по порогу, указанному в параметре *threshold*
        - *func OrderedDithering(srcImgPath string, dstImgPath string, order int) (errF error)* - функция, реализующая алгоритм упорядоченного размытия с использованием матрицы размытия порядка *order*
        - *func FloydErrDithering(srcImgPath string, dstImgPath string, threshold int) (errF error)* - функция, реализующая алгоритм рассеивания ошибок Флойда-Стейнберга с порогом равным *threshold*
    - _**internal/algorithms/dithering/dithering_parallel.go**_
        - *func ThresholdDitheringParallel(srcImgPath string, dstImgPath string, threshold int, threadCount int) (errF error)* - параллельная реализация алгоритма усечения по порогу с порогом равным *threshold* и использованием *threadCount* потоков
        - *func OrderedDitheringParallel(srcImgPath string, dstImgPath string, order int, threadCount int) (errF error)* - параллельная реализация алгоритма упорядоченного размытия с использованием матрицы размытия порядка *order* на *threadCount* числе потоков
        - *func FloydErrDitheringParallel(srcImgPath string, dstImgPath string, threshold int, threadCount int) (errF error)* - параллельная реализация алгоритма рассеивания ошибок Флойда-Стейнберга с порогом равным *threshold* и использованием *threadCount* потоков
- _**internal/algorithms/fourier**_
    - _**internal/algorithms/fourier/fourier.go**_
        - *func DirectTransform(values *[]float64, borders ...int) (result []complex128, errF error)* - функция, реализующая прямое дискретное преобразование Фурье на заданных границах *borders* (должно быть передано два значения, но если не указаны, то берется весь диапазон)
        - *func InverseTransform(values *[]complex128, borders ...int) (result []float64, errF error)* - функция, реализующая обратное дискретное преобразование Фурье на заданных границах *borders* (должно быть передано два значения, но если не указаны, то берется весь диапазон)
    - _**internal/algorithms/fourier/fourier_parallel.go**_
        - *type directRetValues* - структура для передачи результата рассчета прямого ДПФ по каналу
        - *type inverseRetValues* - структура для передачи результата рассчета обратного ДПФ по каналу
        - *func DirectTransformParallel(values *[]float64, threadCount int) (result []complex128, errF error)* - функция-оболочка для параллельного выполнения рассчета прямого ДПФ с использованием *threadCount* потоков
        - *func InverseTransformParallel(values *[]complex128, threadCount int) (result []float64, errF error)* - функция-оболочка для параллельного выполнения рассчета обратного ДПФ с использованием *threadCount* потоков
- _**pkg/analyzer**_
    - _**pkg/analyzer/analyzer.go**_
        - *type Analyzer* - тип, хранящий результаты анализа длительности и предоставляющий методы их обработки
            - *func New(threadCounts []int, testCount int, name ...string) \*Analyzer* - конструктор *Analyzer* с заданным числом потоков и тестов (параметр *name* является необязательным и предназначен для именования тестирования)
            - *func (anl *Analyzer) Analyze(f func(int) error) (errF error)* - функция многократного анализа длительности выполнения функции на заданном множестве потоков
            - *func (anl *Analyzer) SaveXLSX(dstFilePath string) (errF error)* - сохранение результатов тестирования в файл *dstFilePath* в xlsx формате (помимо построения таблицы результатов формируется линейный график для их наглядного сравнения)

### Команды make

- *make run* - запуск выполнения программы
- *make tests* - запуск unit-тестов (существующие директории results и test будут удалены)
- *make testCover* - запуск unit-тестов с флагом -cover (существующие директории results и test будут удалены)