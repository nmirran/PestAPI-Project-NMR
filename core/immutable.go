package core

import "pestapi/model"

func UpdateCommonNameImmutable(id int, newName string) []model.Pest {
	original	:= PestStore.GetAll()
	updated := DeepCopyPests(original)
	
	for i, p := range updated {
		if p.ID == id {
			updated[i].CommonName = newName
			break
		}
	}
	return updated
}

func RemovePestImmutable(id int) []model.Pest {
	original := PestStore.GetAll()
	result := make([]model.Pest, 0, len(original))

	for _, p := range original {
		if p.ID != id {
			result = append(result, p)
		}
	}
	return result
}

func AddPestImmutable(newPest model.Pest) []model.Pest {
	original := PestStore.GetAll()
	result := make([]model.Pest, 0, len(original)+1)
	result = append(result, original...)
	result = append(result, newPest)
	return result
}

func DeepCopyPests(pest []model.Pest) []model.Pest {
	copyList := make([]model.Pest, len(pest))
	for i, p := range pest {
		copyList[i] = model.Pest{
			ID:             p.ID,
			CommonName:     p.CommonName,
			ScientificName: p.ScientificName,
			PestType:       p.PestType,
			AffectedParts:  append([]string{}, p.AffectedParts...),
			Description:    p.Description,
			Symptoms:       append([]string{}, p.Symptoms...),
			ControlMethods: model.ControlMethods{
				Organic:  append([]string{}, p.ControlMethods.Organic...),
				Chemical: append([]string{}, p.ControlMethods.Chemical...),
			},
		}
	}
	return copyList
}

