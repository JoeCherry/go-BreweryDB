package brewerydb

type Style struct {
    Id int
    CategoryId int
    Category Category
    Name string
    ShortName string
    Descrption string
    IbuMin string
    IbuMax string
    AbvMin string
    AbvMax string
    SrmMin string
    SrmMax string
    OgMin string
    FgMin string
    FgMax string
    CreateDate string
    UpdateDate string
}
