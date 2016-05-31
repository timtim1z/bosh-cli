package cmd

import (
	"strings"

	boshdir "github.com/cloudfoundry/bosh-init/director"
	boshui "github.com/cloudfoundry/bosh-init/ui"
	boshtbl "github.com/cloudfoundry/bosh-init/ui/table"
)

type LocksCmd struct {
	ui       boshui.UI
	director boshdir.Director
}

func NewLocksCmd(ui boshui.UI, director boshdir.Director) LocksCmd {
	return LocksCmd{ui: ui, director: director}
}

func (c LocksCmd) Run() error {
	locks, err := c.director.Locks()
	if err != nil {
		return err
	}

	table := boshtbl.Table{
		Content: "locks",
		Header:  []string{"Type", "Resource", "Expires at"},
		SortBy:  []boshtbl.ColumnSort{{Column: 2, Asc: true}},
	}

	for _, l := range locks {
		table.Rows = append(table.Rows, []boshtbl.Value{
			boshtbl.ValueString{l.Type},
			boshtbl.ValueString{strings.Join(l.Resource, ":")},
			boshtbl.ValueTime{l.ExpiresAt},
		})
	}

	c.ui.PrintTable(table)

	return nil
}